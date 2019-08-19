package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

const (
	wereadChapterURLPattern = "https://weread.qq.com/wrpage/wechat/search/read/chapter?uid=%d&bookId=%d"
	wereadCoverURLPattern   = "https://weread.qq.com/wrpage/wechat/search/book/%d?ref=mp"
	wereadRequestUserAgent  = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36 QBCore/3.53.1159.400 QQBrowser/9.0.2524.400 Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 MicroMessenger/6.5.2.501 NetType/WIFI WindowsWechat"
)

//Book weread book
type Book struct {
	BookID       int
	Title        string
	Author       string
	Chapters     []*Chapter
	CoverImgData []byte
	CoverImgExt  string
}

//Chapter chapter
type Chapter struct {
	Data *ChapterData
}

//ChapterData chapter data
type ChapterData struct {
	Content string
	Title   string
	Idx     int
	Uid     int
	NextUid int
	LastUid int
	Price   int
}

//ContentHTML chapter content html
func (chData *ChapterData) ContentHTML() template.HTML {
	return template.HTML(chData.Content)
}

//String chapter to string
func (ch *Chapter) String() string {
	return fmt.Sprintf("%#v", ch.Data)
}

//NewBook create a book
func NewBook(bookID int) *Book {
	return &Book{BookID: bookID}
}

//WereadBook load data to a weread book
func (book *Book) WereadBook(cookie string) error {
	chapters := make([]*Chapter, 0)
	uID := 1
	for {
		chapter, err := book.WereadChapter(uID, cookie)
		if err != nil {
			return err
		}
		chapters = append(chapters, chapter)
		if chapter.Data.NextUid < 1 {
			break
		}
		uID++
	}
	book.Chapters = chapters
	err := book.WereadMeta(cookie)
	if err != nil {
		return err
	}
	return nil
}

//WereadMeta load meta info from a book
func (book *Book) WereadMeta(cookie string) error {
	coverURL := fmt.Sprintf(wereadCoverURLPattern, book.BookID)
	fmt.Printf("weread cover:%s\n", coverURL)
	content, err := httpGet(coverURL, cookie)
	if err != nil {
		return err
	}
	html := string(content)
	var author string
	var coverImgURL string
	html, author, err = extractInfo(html, "wr_bookInfoHeader_author\">", "</p>")
	if err != nil {
		return err
	}
	book.Author = author
	html, coverImgURL, err = extractInfo(html, "config.bookCover = '", "';")
	if err != nil {
		return err
	}
	fmt.Printf("weread cover img:%s\n", coverImgURL)
	coverImgData, err := httpGet(coverImgURL, "")
	if err != nil {
		return err
	}
	book.CoverImgData = coverImgData
	book.CoverImgExt = filepath.Ext(coverImgURL)

	_, title, err := extractInfo(html, "config.bookTitle = '", "';")
	if err != nil {
		return err
	}
	book.Title = title

	return nil
}

//WereadChapter get weread chapter content from a book
//Error Message like: {"statusCode":200,"method":"GET","result":{"errCode":-2012,"humanMessage":"登录超时"}}
func (book *Book) WereadChapter(uID int, cookie string) (*Chapter, error) {
	chapterURL := fmt.Sprintf(wereadChapterURLPattern, uID, book.BookID)
	fmt.Printf("weread chapter:%s\n", chapterURL)
	content, err := httpGet(chapterURL, cookie)
	if err != nil {
		return nil, err
	}
	var chapter Chapter
	err = json.Unmarshal(content, &chapter)
	if err != nil {
		return nil, err
	}
	if chapter.Data == nil {
		return nil, fmt.Errorf("unexpected chapter data:%s", string(content))
	}
	return &chapter, nil
}

func extractInfo(info, start, end string) (string, string, error) {
	idx := strings.Index(info, start)
	if idx < 1 {
		return info, "", fmt.Errorf("failed to find info, start Idx:%d", idx)
	}
	info = info[idx+len(start):]
	idx = strings.Index(info, end)
	if idx < 1 {
		return info, "", fmt.Errorf("failed to find info, end Idx:%d", idx)
	}
	return info, info[:idx], nil
}

func httpGet(url, cookie string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if cookie != "" {
		req.Header.Set("Cookie", cookie)
	}
	req.Header.Set("User-Agent", wereadRequestUserAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
