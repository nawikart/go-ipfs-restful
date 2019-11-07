# API DEV
<br />

## Deploy to docker

> make sure you already install docker and docker compose

- https://docs.docker.com/install/
- https://docs.docker.com/compose/install/ 
<br />

> then run the following command:

``` bash
$ dockerd
$ docker-compose up
```
<br><br>

## Endpoints

<details>
<summary>ADD</summary>

###### *Add single/multiple string and or file to IPFS*

> POST: http://localhost:8081/ipfs/add

``` bash
# headers:
Content-Type: multipart/form-data
```

``` bash
# sample body

# single data 
file: sample.jpg # type file

# single data 
arg: "Lorem ipsum..." # type string

# multiple data
file: sample.jpg # type file
file: sample2.jpg # type file

# multiple data
file: sample.jpg # type file
file: sample2.jpg # type file
arg: "Lorem ipsum..." # type string
arg: "Sample 2 Lorem ipsum..." # type string

```

``` bash
# sample success response:
[
    {
        "status": "success",
        "data": {
            "data": "sample.jpg",
            "hash": "QmWaV5TVkGAR4infA3zP2mqTjwBUHcZy6YrTw4eyvTxzej"
        }
    },
    {
        "status": "success",
        "data": {
            "data": "sample2.jpg",
            "hash": "QmWaV5TVkGAR4infA3zP2mqTjwBUHcZy6YrTw4eyvTxuyt"
        }
    },
    {
        "status": "success",
        "data": {
            "data": "Lorem ipsum...",
            "hash": "QmXy3FjCoBGcFjcKwCdYZDsqQgmsscQ79LNtwZfNkp4L7g"
        }
    },
    {
        "status": "success",
        "data": {
            "data": "Sample 2 Lorem ipsum...",
            "hash": "QmRhkA7aXDhTeGDA959dGLxBKLHKbCc9FuWtevKxwSTqL3"
        }
    }
]
```

``` bash
# sample failed response:
{
    "status": "failed",
    "error_msg": "multipart: NextPart: EOF"
}

# OR

[
    {
        "status": "failed",
        "error_msg": "Post http://ipfs_host/5001/api/v0/add?: read tcp 172.18.0.2:53164->202.3.219.211:80: read: connection reset by peer",
        "data": {
            "data": "dog.jpeg"
        }
    },
    {
        "status": "failed",
        "error_msg": "Post http://ipfs_host/5001/api/v0/add?: read tcp 172.18.0.2:53186->202.3.219.211:80: read: connection reset by peer",
        "data": {
            "data": "abc"
        }
    }
]
```
</details>
<hr /><br>


<details>
<summary>ADD DIRECTORY</summary>

###### *adds a directory recursively with all of the files under it*

> POST: http://localhost:8081/ipfs/adddir

``` bash
# headers:
Content-Type: application/json
```

``` bash
# sample body
{
	"directory": "/home/sample/foo"
}
```

``` bash
# sample success response:
{
    "status": "success",
    "data": {
        "data": "/home/sample/foo",
        "hash": "QmdwSiWRYypAv4hgWUAcNfQcnkTQ72C37Pgck92iQdtphs"
    }
}
```

``` bash
# sample failed response:
{
    "status": "success",
    "error_msg": "lstat /home/sample/foo: no such file or directory",
    "data": {
        "data": "/home/sample/foo"
    }
}
```
</details>
<hr /><br>

<details>
<summary>CAT</summary>

###### *Cat the content at the given path.*

> GET: http://localhost:8081/ipfs/cat/:hash

``` bash
# success response accesed via browser will showing the content string/image/pdf. If the content are .zip then will auto be downloaded.
```

``` bash
# sample failed response:
{
"status": "failed",
"error_msg": "cat: invalid path \"QmWkTKsj33qSEV4Yghba6fBjtiLX76K77bvh6S3qFemFACx\": selected encoding not supported"
}
```
</details>
<hr /><br>


<details>
<summary>LIST</summary>

###### *List entries at the given path.*

> GET: http://localhost:8081/ipfs/list

``` bash
# sample success response:
{
    "status": "success",
    "data": [
        {
            "Hash": "QmeBpzHngbHes9hoPjfDCmpNHGztkmZFRX4Yp9ftKcXZDN",
            "Name": "bar",
            "Size": 0,
            "Type": 1
        },
        {
            "Hash": "QmSbAeaYBpavHpDuzV9hugbLLS1EipE3BwWHKZgtfeH5nv",
            "Name": "baz",
            "Size": 4,
            "Type": 2
        },
        {
            "Hash": "Qmd286K6pohQcTKYqnS1YhWrCiS4gz7Xi34sdwMe9USZ7u",
            "Name": "cat.jpg",
            "Size": 443230,
            "Type": 2
        }
    ]
}
```

``` bash
# sample failed response:
{
    "status": "failed",
    "error_msg": "ls: invalid path \"QmdwSiWRYypAv4hgWUAcNfQcnkTQ72C37Pgck92iQdtphsx\": selected encoding not supported"
}
```
</details>
<hr /><br>

<details>
<summary>REFS</summary>

###### *Lists links (references) from an object.*

> POST: http://localhost:8081/ipfs/refs

``` bash
# headers:
Content-Type: application/json
```

``` bash
# sample body
{
	"hash": "QmdwSiWRYypAv4hgWUAcNfQcnkTQ72C37Pgck92iQdtphs",
	"recursive": true
}
```

``` bash
# sample success response:
{
    "status": "success",
    "data": [
        "QmeBpzHngbHes9hoPjfDCmpNHGztkmZFRX4Yp9ftKcXZDN",
        "QmWLdkp93sNxGRjnFHPaYg8tCQ35NBY3XPn6KiETd3Z4WR",
        "QmSbAeaYBpavHpDuzV9hugbLLS1EipE3BwWHKZgtfeH5nv",
        "Qmd286K6pohQcTKYqnS1YhWrCiS4gz7Xi34sdwMe9USZ7u",
        "QmPEKipMh6LsXzvtLxunSPP7ZsBM8y9xQ2SQQwBXy5UY6e",
        "QmT8onRUfPgvkoPMdMvCHPYxh98iKCfFkBYM1ufYpnkHJn"
    ]
}
```

``` bash
# sample failed response:
{
    "status": "failed",
    "error_msg": "refs: invalid path \"QmdwSiWRYypAv4hgWUAcNfQcnkTQ72C37Pgck92iQdtphsx\": selected encoding not supported"
}
```
</details>
<hr /><br>


<details>
<summary>PIN</summary>

###### *Pin the given path.*

> GET: http://localhost:8081/ipfs/pin/:hash

``` bash
# sample success response:
{
    "status": "success"
}
```

``` bash
# sample failed response:
{
    "status": "failed",
    "error_msg": "pin/add: invalid path \"QmQnvUyGs6hnhEk4S57uSWvpD9JC9ZaBNuKQoU6JZtiXNYx\": selected encoding not supported"
}
```
</details>
<hr /><br>



<details>
<summary>UNPIN</summary>

###### *Unpin the given path.*

> GET: http://localhost:8081/ipfs/unpin/:hash

``` bash
# sample success response:
{
    "status": "success"
}
```

``` bash
# sample failed response:
{
    "status": "failed",
    "error_msg": "pin/rm: invalid path \"QmQnvUyGs6hnhEk4S57uSWvpD9JC9ZaBNuKQoU6JZtiXNYx\": selected encoding not supported"
}
```
</details>
<hr /><br>


<details>
<summary>PINS</summary>

###### *Pins returns a map of the pin hashes to their info (currently just the pin type, one of DirectPin, RecursivePin, or IndirectPin.*

> GET: http://localhost:8081/ipfs/pins

``` bash
# sample success response:
{
    "status": "success",
    "data": {
        "QmNM52yu9pdXzYDZAiESQzGdxvS7Aex7rhY2MHkAy8mrG9": {
            "Type": "indirect"
        },
        "QmNM5Ly3hX8UVXFvDvVGyqkPxgZjXuW9xC9Dy4hRqwNCyV": {
            "Type": "recursive"
        },
        "QmNM8AKoCDCcYnAQ6j1JQHiaP7Jr7kfiostnuY3SMjHddU": {
            "Type": "indirect"
        }
    }
}
```
</details>
<hr /><br>


<details>
<summary>NAME PUBLISH</summary>

###### *Publish IPNS names.*

> POST: http://localhost:8081/ipfs/name/publish

``` bash
# headers:
Content-Type: application/json
```

``` bash
# sample body
{
	"contenthash": "Qmbu7x6gJbsKDcseQv66pSbUcAA3Au6f7MfTYVXwvBxN2K",
	"key": "mykey",
	"lifetime": 3600,
	"ttl": 3600,
	"resolve": false
}
```

``` bash
# sample success response:
{
    "status": "success",
    "data": {
        "name": "QmQdg7mmxnQ3xvybjk1xGpBHU5QMsQKEbUeCgxrBg1pAKK",
        "value": "/ipfs/Qmbu7x6gJbsKDcseQv66pSbUcAA3Au6f7MfTYVXwvBxN2K"
    }
}
```

``` bash
# sample failed response:
{
    "status": "failed",
    "error_msg": "name/publish: failed to find any peer in table"
}

# OR

{
    "status": "failed",
    "error_msg": "name/publish: expired record"
}
```
</details>
<hr /><br>


<details>
<summary>NAME RESOLVE</summary>

###### *Resolve gets resolves the string provided to an /ipns/[name]. If asked to resolve an empty string, resolve instead resolves the node's own /ipns value.*

> GET: http://localhost:8081/ipfs/name/resolve/:hash

``` bash
# sample success response:
{
    "status": "success",
    "data": "/ipfs/QmWe8FaHMRb3fGLqyyvxSpkAicMqGnLkmr1fMrr4Ep8SNT"
}
```

``` bash
# sample failed response:
{
    "status": "failed",
    "error_msg": "name/resolve: context deadline exceeded"
}

# OR 

{
    "status": "failed",
    "error_msg": "name/resolve: failed to dial QmSiTko9JZyabH56y2fussEt1A5oDqsFXB3CkvAqraFryz: no addresses"
}
```
</details>
<hr /><br>