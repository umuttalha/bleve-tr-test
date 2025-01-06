//  Copyright (c) 2018 Couchbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tr

import (
	"fmt"

	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"

	"github.com/blevesearch/snowballstem"
	"github.com/blevesearch/snowballstem/turkish"
)

const SnowballStemmerName = "stemmer_tr_snowball"

type TurkishStemmerFilter struct {
}

func NewTurkishStemmerFilter() *TurkishStemmerFilter {
	return &TurkishStemmerFilter{}
}

func (s *TurkishStemmerFilter) Filter(input analysis.TokenStream) analysis.TokenStream {

	for _, token := range input {

		// buradan string(token.Term) alıp grpc ye yolla

		fmt.Print(string(token.Term))
		fmt.Print("\n")

		env := snowballstem.NewEnv(string(token.Term))
		turkish.Stem(env)

		// buradan outputu alıp env.SetCurrent("asd") ile set et

		fmt.Print(env.Current())
		fmt.Print("\n")

		token.Term = []byte(env.Current())
	}

	return input
}

func TurkishStemmerFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenFilter, error) {
	return NewTurkishStemmerFilter(), nil
}

func init() {
	registry.RegisterTokenFilter(SnowballStemmerName, TurkishStemmerFilterConstructor)
}
