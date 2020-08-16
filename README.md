# img2ascii
 Simple term ascii image encoder fully tested.

 * **Test**
 update golden files
 ```sh
 $ go go test ./... -update
 ``` 

```sh
 $ go go test ./... -cover -v
 ``` 

 * **Build**
 ```sh
 $ go build -v
 ```

 * **Usage**

  ```sh
$ ./img2ascii -h
Usage of ./img2ascii:
  -f string
        -f <filename> (default "sample.jpg")
  -s uint
        -s <scale> (default 70)
```

```sh
$ ./img2ascii -f encoder/testdata/tux.jpg -s 60
```

<img src="results/enctux.PNG" width="60%">

```sh
$ ./img2ascii -f encoder/testdata/julio.jpg -s 60
```

<img src="results/encjulio.PNG">


