/*
 * Minimalist Object Storage, (C) 2014,2015 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"github.com/minio-io/cli"
	"github.com/minio-io/mc/pkg/s3"
)

// doMakeBucketCmd creates a new bucket
func doMakeBucketCmd(c *cli.Context) {
	urlStr, err := parseURL(c)
	if err != nil {
		fatal(err.Error())
	}

	bucket, err := url2Bucket(urlStr)
	if err != nil {
		fatal(err.Error())
	}

	s3c, err := getNewClient(c.GlobalBool("debug"), urlStr)
	if err != nil {
		fatal(err.Error())
	}

	if !s3.IsValidBucketName(bucket) {
		fatal(errInvalidbucket.Error())
	}

	err = s3c.PutBucket(bucket)
	if err != nil {
		fatal(err.Error())
	}
}
