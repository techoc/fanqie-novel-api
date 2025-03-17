package main

import (
	"fmt"
	"testing"

	"github.com/techoc/fanqie-novel-api/pkg/fanqie"
)

func TestGetContentByChapterIdV2(t *testing.T) {
	chapter := fanqie.GetContentByChapterIdV2(7461737987731112510)
	fmt.Println(chapter.Content)
}

func TestGetNewCookie(t *testing.T) {
	fanqie.GetNewCookie()
}
