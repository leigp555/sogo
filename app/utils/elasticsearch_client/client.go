package elasticsearch_client

import (
	es "github.com/elastic/go-elasticsearch/v7"
	"sogo/app/global/my_errors"
	"sogo/app/global/variable"
)

func NewElasticsearchClient() *es.Client {
	addr := variable.Config.GetString("elasticsearch.addr")
	username := variable.Config.GetString("elasticsearch.username")
	password := variable.Config.GetString("elasticsearch.password")
	client, err := es.NewClient(es.Config{
		Addresses: []string{addr},
		Username:  username,
		Password:  password,
	})
	if err != nil {
		panic(my_errors.ErrorsElasticSearchInitFail + err.Error())
	}
	return client
}
