package main

import (
	"github.com/GabrielLuizSF/go_url/pkg/system/analyzer"
	"github.com/GabrielLuizSF/go_url/pkg/version"
)

// Copyright 2023 gabrielluizsf. All rights reserved
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

func main(){
	version.PrintVersion();
    analyzer.OSAnalyzer();
}


