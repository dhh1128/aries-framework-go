// +build js,wasm

/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/piprate/json-gold/ld"
)

const jsonldContextPrefix = "data/context"

func createLDPBBS2020DocumentLoader() ld.DocumentLoader {
	loader := ld.NewCachingDocumentLoader(ld.NewRFC7324CachingDocumentLoader(&http.Client{}))

	addJSONLDCachedContextFromFile(loader,
		"https://www.w3.org/2018/credentials/v1", "vc.jsonld")

	addJSONLDCachedContextFromFile(loader,
		"https://w3c-ccg.github.io/ldp-bbs2020/context/v1", "ldp-bbs2020.jsonld")

	addJSONLDCachedContextFromFile(loader,
		"https://w3id.org/security/v1", "security_v1.jsonld")

	addJSONLDCachedContextFromFile(loader,
		"https://w3id.org/security/v2", "security_v2.jsonld")

	addJSONLDCachedContextFromFile(loader,
		"https://w3id.org/citizenship/v1", "citizenship.jsonld")

	return loader
}

func addJSONLDCachedContextFromFile(loader *ld.CachingDocumentLoader, contextURL, contextFile string) {
	contextPath := filepath.Clean(filepath.Join(jsonldContextPrefix, contextFile))

	contextContent, err := ioutil.ReadFile(contextPath) //nolint:gosec
	if err != nil {
		panic(err)
	}

	addJSONLDCachedContext(loader, contextURL, string(contextContent))
}

func addJSONLDCachedContext(loader *ld.CachingDocumentLoader, contextURL, contextContent string) {
	reader, err := ld.DocumentFromReader(strings.NewReader(contextContent))
	if err != nil {
		panic(err)
	}

	loader.AddDocument(contextURL, reader)
}
