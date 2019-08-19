package main

import "html/template"

const (
	//OpfTmpXML opf template xml content
	OpfTmpXML = `
	{{.XmlHeader}}
	<package unique-identifier="uid">
	<metadata>
	  <dc-metadata 
		xmlns:dc="http://purl.org/metadata/dublin_core"
		xmlns:oebpackage="http://openebook.org/namespaces/oeb-package/1.0/">
		<dc:Title>{{.Book.Title}}</dc:Title>
		<dc:Publisher>Weread</dc:Publisher>
		<dc:Creator>{{.Book.Author}}</dc:Creator>
		<dc:Producer>kindlegen</dc:Producer>
		<dc:Language>en-us</dc:Language>
		</dc-metadata>
		<x-metadata>
		  <output encoding="UTF-8" content-type="text/x-oeb1-document"></output>	
		  <EmbeddedCover>cover{{.Book.CoverImgExt}}</EmbeddedCover>
		</x-metadata>
	</metadata>
	
	<manifest>
	   <item id="book" media-type="text/x-oeb1-document" href="book.html"></item>
	   <item id="toc" media-type="application/x-dtbncx+xml" href="book.ncx"/>
	   <item id="bookcover" media-type="image/png" href="cover{{.Book.CoverImgExt}}"></item>
	</manifest>
	
	<spine toc="toc" pageList></spine>
	</package>
	`

	//NcxTmpXML ncx tmplate XML
	NcxTmpXML = `
	{{.XmlHeader}}
	<!DOCTYPE ncx PUBLIC "-//NISO//DTD ncx 2005-1//EN" "http://www.daisy.org/z3986/2005/ncx-2005-1.dtd">
	<ncx xmlns="http://www.daisy.org/z3986/2005/ncx/"  version="2005-1" xml:lang="en-US">
	<head>
	<meta name="dtb:uid" content="uid"/>
	<meta name="dtb:depth" content="1"/>
	<meta name="dtb:totalPageCount" content="0"/>
	<meta name="dtb:maxPageNumber" content="0"/>
	</head>
	<docTitle><text>{{.Book.Title}}</text></docTitle>
	<docAuthor><text>{{.Book.Author}}</text></docAuthor>
	<navMap>
		{{range $idx, $chapter := .Book.Chapters}}
		<navPoint id="chapter-{{$chapter.Data.Idx}}" playOrder="{{$chapter.Data.Idx}}">
		<navLabel><text>{{$chapter.Data.Title}}</text></navLabel>
		<content src="book.html#chapter-{{$chapter.Data.Idx}}"/>
		</navPoint>
		{{end}}
	</navMap>
	</ncx>
	`
	//BookTmpHTML book template html
	BookTmpHTML = `
	<!DOCTYPE html>
	<html lang="zh-cmn">
		<head>
			<meta http-equiv="content-type" content="text/html; charset=UTF-8">
			<meta http-equiv="X-UA-Compatible" content="IE=Edge"/>
			<meta name="renderer" content="webkit"/>
			<meta name="viewport" content="viewport-fit=cover,width=device-width,initial-scale=1,maximum-scale=1,user-scalable=0"/>
			<meta name="format-detection" content="telephone=no"/>
			<title>{{.Book.Title}}</title>
			<style>
				::-webkit-scrollbar {
					width: 12px !important;
					height: 12px !important;
				}
	
				::-webkit-scrollbar-track:vertical {
				}
	
				::-webkit-scrollbar-thumb:vertical {
					background-color: rgba(136, 141, 152, 0.5) !important;
					border-radius: 10px !important;
					background-clip: content-box !important;
					border: 2px solid transparent !important;
				}
	
				::-webkit-scrollbar-track:horizontal {
				}
	
				::-webkit-scrollbar-thumb:horizontal {
					background-color: rgba(136, 141, 152, 0.5) !important;
					border-radius: 10px !important;
					background-clip: content-box !important;
					border: 2px solid transparent !important;
				}
	
				::-webkit-resizer {
					display: none !important;
				}
			</style>
		</head>
		<body>
			<div class="book">
			{{range $idx, $chapter := .Book.Chapters}}
				<div class="chapter">
					<a id="chapter-{{$chapter.Data.Idx}}"></a>
					<div class="chapter_title">{{$chapter.Data.Title}}</div>
					<div class="chapter_content">
						{{$chapter.Data.ContentHTML}}
					</div>
				</div>
			{{end}}
			</div>
		</body>
	</html>
	`
)

var (
	opfTmp  = template.Must(template.New("opf").Parse(OpfTmpXML))
	ncxTmp  = template.Must(template.New("ncx").Parse(NcxTmpXML))
	bookTmp = template.Must(template.New("book").Parse(BookTmpHTML))
)
