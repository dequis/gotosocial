/*
   GoToSocial
   Copyright (C) 2021-2022 GoToSocial Authors admin@gotosocial.org

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package storage

import (
	"context"
	"mime"
	"net/url"
	"path"
	"time"

	"codeberg.org/gruf/go-store/v2/kv"
	"codeberg.org/gruf/go-store/v2/storage"
)

type S3 struct {
	Proxy   bool
	Bucket  string
	Storage *storage.S3Storage
	*kv.KVStore
}

func (s *S3) URL(ctx context.Context, key string) *url.URL {
	if s.Proxy {
		return nil
	}

	// it's safe to ignore the error here, as we just fall back to fetching the file if URL request fails
	url, _ := s.Storage.Client().PresignedGetObject(ctx, s.Bucket, key, time.Hour, url.Values{
		"response-content-type": []string{mime.TypeByExtension(path.Ext(key))},
	})

	return url
}
