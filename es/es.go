package es

import (
	"context"
	"log"
	"os"

	"encoding/json"

	"github.com/olivere/elastic"
)

type UserInfo struct {
	UID int64 `json:"uid"`
}

func Search() {
	log.Print("hello world")
	ctx := context.Background()
	client, err := elastic.NewClient(
		elastic.SetURL("http://es.dsj.inkept.cn:80"),
		elastic.SetSniff(false),
		elastic.SetTraceLog(log.New(os.Stderr, "", log.LstdFlags)),
	)
	if err != nil {
		log.Printf("new client err, %s", err)
		return
	}
	info, code, err := client.Ping("http://es.dsj.inkept.cn:80").Do(ctx)
	if err != nil {
		log.Printf("ping client err, %s", err)
		return
	}
	log.Print("info, code", info, code)
	boolQuery := elastic.NewBoolQuery()
	anQuery := elastic.NewMatchQuery("appname", "ms")
	indexQuery := elastic.NewMatchQuery("index_type", "user_profile")
	rangeQuery := elastic.NewRangeQuery("last_login_date").Gte("20210127")
	boolQuery.Must(anQuery, indexQuery, rangeQuery)
	ss := elastic.NewSearchSource().Query(boolQuery).Size(20)
	ssSource, _ := ss.Source()
	ssJson, err := json.Marshal(ssSource)
	log.Print(string(ssJson))
	res, err := client.Search().Index("app_uid").SearchSource(ss).Do(ctx)
	if err != nil {
		log.Printf("search error, %s", err)
		return
	}
	var uids []int64
	var item UserInfo
	for _, hit := range res.Hits.Hits {
		err = json.Unmarshal(*hit.Source, &item)
		uids = append(uids, item.UID)
	}
	log.Printf("%#v", len(uids))
}

func SearchWithScroll() {
	log.Print("search")
	ctx := context.Background()
	client, err := elastic.NewClient(
		elastic.SetURL("http://es.dsj.inkept.cn:80"),
		elastic.SetSniff(false),
		elastic.SetTraceLog(log.New(os.Stderr, "", log.LstdFlags)),
	)
	if err != nil {
		log.Printf("new client err, %s", err)
		return
	}
	anq := elastic.NewMatchQuery("appname", "ms")
	itq := elastic.NewMatchQuery("index_type", "user_profile")
	rgq := elastic.NewRangeQuery("last_login_date").Gte("20210127")
	bq := elastic.NewBoolQuery().Must(anq, itq, rgq)
	ss := elastic.NewSearchSource().Query(bq)

	var scrollID string
	res, err := client.Scroll().Index("app_uid").SearchSource(ss).Scroll("5m").Size(200).Do(ctx)
	if err != nil {
		log.Printf("scroll err, %s", err)
	}
	scrollID = res.ScrollId
	var uids []int64
	var item UserInfo
	for _, hit := range res.Hits.Hits {
		err = json.Unmarshal(*hit.Source, &item)
		uids = append(uids, item.UID)
	}
	log.Printf("scroll id: %s", res.ScrollId)
	log.Printf("uids:%#v", uids)
	for len(scrollID) > 0 {
		res, err = client.Scroll().Scroll("1m").Index("app_uid").ScrollId(scrollID).Size(200).Do(ctx)
		if err != nil {
			log.Printf("scroll continue err, %s", err)
			break
		}
		scrollID = res.ScrollId
		for _, hit := range res.Hits.Hits {
			err = json.Unmarshal(*hit.Source, &item)
			uids = append(uids, item.UID)
		}
		log.Printf("scroll id: %s", res.ScrollId)
		log.Printf("uids:%#v", len(uids))
	}
	log.Printf("scroll search done: %d", len(uids))
}
