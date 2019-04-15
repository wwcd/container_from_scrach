Using alpine as rootfs

    ❯ docker inspect alpine:3.5 -f "{{ .GraphDriver.Data.UpperDir }}"
    /var/lib/docker/overlay2/254f7f0509dfae6f4a11c1f11149ce1367ccc3415f5727ce2e75d5891f3de68e/diff

    ❯ mkdir rootfs/{upperlayer,workdir,mountedfs}

    ❯ cd rootfs && mount -t overlay -o lowerdir=/var/lib/docker/overlay2/254f7f0509dfae6f4a11c1f11149ce1367ccc3415f5727ce2e75d5891f3de68e/diff,upperdir=upperlayer,workdir=workdir overlay mountedfs

Run simplecontainer

    ❯ sudo /opt/go/bin/go run main.go run sh
    Container PID is  13872
    Container init PID is  1
    / #
