# log-monitor

### 日志监控服务

* 数据上报 ElasticSearch
* 定时监控告警

### 部署 es

* https://www.elastic.co/cn/downloads/elasticsearch
* curl -OL https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-7.15.2-linux-x86_64.tar.gz
* tar -zxvf elasticsearch-7.15.2-linux-x86_64.tar.gz
* 配置文件
    * xpack.security.enabled: true
    * xpack.license.self_generated.type: basic
    * xpack.security.transport.ssl.enabled: true
* ./elasticsearch -d
* adduser elasticsearch
* passwd elasticsearch
* chown -R elasticsearch elasticsearch-7.15.2
* su elasticsearch
* ./elasticsearch -d
* curl http://elastic:elasticsearch@127.0.0.1:9200
* ./bin/elasticsearch-setup-passwords interactive

### 磁盘扩容

* fdisk -lu
* df -Th
* https://help.aliyun.com/document_detail/25426.html?spm=a2c4g.11186623.2.14.4b7d32fdD7SmRB
* fdisk -u /dev/vdb
    * p,n,p,换行，换行，换行，p,w
* mkfs -t ext4 /dev/vdb1
* cp /etc/fstab /etc/fstab.bak
* echo `blkid /dev/vdd1 | awk '{print $2}' | sed 's/\"//g'` /vdd ext4 defaults 0 0 >> /etc/fstab
* mount -a