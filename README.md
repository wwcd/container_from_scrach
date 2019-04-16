Using alpine as rootfs

    ❯ LOWERLAYER=$(docker inspect alpine:3.5 -f "{{ .GraphDriver.Data.UpperDir }}")

    ❯ mkdir rootfs/{upperlayer,workdir,mountedfs}

    ❯ cd rootfs && mount -t overlay -o lowerdir=${LOWERLAYER},upperdir=upperlayer,workdir=workdir overlay mountedfs

Run simplecontainer

    ❯ go run main.go run sh
    Container PID is  13872
    Container init PID is  1
    / #
