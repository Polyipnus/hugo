// Copyright 2018 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package privacy

import (
	"testing"

	"github.com/gohugoio/hugo/config"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestDecodeConfigFromTOML(t *testing.T) {
	assert := require.New(t)

	tomlConfig := `

someOtherValue = "foo"

[privacy]
[privacy.disqus]
disable = true
[privacy.googleAnalytics]
disable = true
[privacy.instagram]
disable = true
[privacy.speakerDeck]
disable = true
[privacy.tweet]
disable = true
[privacy.vimeo]
disable = true
[privacy.youtube]
disable = true
noCookie = true
`
	cfg, err := config.FromConfigString(tomlConfig, "toml")
	assert.NoError(err)

	pc, err := DecodeConfig(cfg)
	assert.NoError(err)
	assert.NotNil(pc)

	assert.True(pc.Disqus.Disable)
	assert.True(pc.GoogleAnalytics.Disable)
	assert.True(pc.Instagram.Disable)
	assert.True(pc.SpeakerDeck.Disable)
	assert.True(pc.Tweet.Disable)
	assert.True(pc.Vimeo.Disable)

	assert.True(pc.YouTube.NoCookie)
	assert.True(pc.YouTube.Disable)
}

func TestDecodeConfigFromTOMLCaseInsensitive(t *testing.T) {
	assert := require.New(t)

	tomlConfig := `

someOtherValue = "foo"

[Privacy]
[Privacy.YouTube]
NoCOOKIE = true
`
	cfg, err := config.FromConfigString(tomlConfig, "toml")
	assert.NoError(err)

	pc, err := DecodeConfig(cfg)
	assert.NoError(err)
	assert.NotNil(pc)
	assert.True(pc.YouTube.NoCookie)
}

func TestDecodeConfigDefault(t *testing.T) {
	assert := require.New(t)

	pc, err := DecodeConfig(viper.New())
	assert.NoError(err)
	assert.NotNil(pc)
	assert.False(pc.YouTube.NoCookie)
}
