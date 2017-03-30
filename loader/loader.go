package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type OpenCVData struct {
	Cascade Cascade `xml:"cascade"`
}

type Cascade struct {
	Height 			int 		`xml:"height"`
	Width 			int			`xml:"width"`
	StageParams 	StageParam	`xml:"stageParams"`
	FeatureParams 	FeatureParam`xml:"featureParams"`
	StageNum 		int			`xml:"stageNum"`
	Stages 			Stages		`xml:"stages"`
	Features		FeatureSet	`xml:"features"`
}

type StageParam struct {
	MaxWeakCount int `xml:"maxWeakCount"`
}

type FeatureParam struct {
	MaxCatCount int `xml:"maxCatCount"`
}

type Stages struct {
	Stage []Stage `xml:"_"`
}

type Stage struct {
	MaxWeakCount 	int 			`xml:"maxWeakCount"`
	StageThreshold 	string 			`xml:"stageThreshold"`
	WeakClassifiers []ClassifierSet `xml:"weakClassifiers"`
}

type ClassifierSet struct {
	Classifiers []Classifier `xml:"_"`
}

type Classifier struct {
	InternalNodes 	string `xml:"internalNodes"`
	LeafValues 		string `xml:"leafValues"`
}

type FeatureSet struct {
	Features []Feature `xml:"_"`
}

type Feature struct {
	Rects TileSet `xml:"rects"`
}

type TileSet struct {
	Tiles []string `xml:"_"`
}

type Query struct {
	Series Show
	// Have to specify where to find episodes since this
	// doesn't match the xml tags of the data that needs to go into it
	EpisodeList []Episode `xml:"Episode"`
}

type Show struct {
	// Have to specify where to find the series title since
	// the field of this struct doesn't match the xml tag
	Title    string `xml:"SeriesName"`
	SeriesID int
	Keywords map[string]bool
}

type Episode struct {
	SeasonNumber  int
	EpisodeNumber int
	EpisodeName   string
	FirstAired    string
}


func (s Show) String() string {
	return fmt.Sprintf("%s - %d", s.Title, s.SeriesID)
}

func (e Episode) String() string {
	return fmt.Sprintf("S%02dE%02d - %s - %s", e.SeasonNumber, e.EpisodeNumber, e.EpisodeName, e.FirstAired)
}

func main() {
	wd, _ := os.Getwd()
	// load haarcascade_frontalface_default.xml
	xmlFile, err := os.Open(wd + "/loader/haarcascade_frontalface_default.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	cascadeFile, _ := ioutil.ReadAll(xmlFile)

	var cvData OpenCVData
	xml.Unmarshal(cascadeFile, &cvData)

	fmt.Println(cvData.Cascade.Height)
	//for _, episode := range q.EpisodeList {
	//	fmt.Printf("\t%s\n", episode)
	//}
}