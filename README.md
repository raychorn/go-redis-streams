# go-redis-streams
Demo using messaging with Redis Streams in Golang

**Using messengerin practice with Redis streams and Golang**

[![YouTube Video Explanation](http://img.youtube.com/vi/Kc-tcrP0c10/0.jpg)](http://www.youtube.com/watch?v=Kc-tcrP0c10 "Utilizando mensageria na pratica com Redis streams e Golang")

# Redis

**What is it?**

Redis is an OpenSource non-relational database, which has key-value storage within its structure. Redis has strategies for storing data in memory and on disk, ensuring fast response and data persistence. Redis's main use cases include caching, session management, PUB/SUB.

# Redis Streams para Mensageria (ou Messaging)

![Design of flow](/media/flow.png)

**Positive Points**

Supports Topicos and Queues
Persistence on disk (through RDB files)
High availability (with Clusterizacao)
High Throughput
Allows Reprocessing
Owns Consumer Groups
Latencia minima
No need for zookeper
It takes up far fewer resources in relation to (Kafka/RabbitMQ)

**Negatives**

Does not guarantee delivery order (yet)
Msgs processed with error does not return for redistribution


# Links

https://www.youtube.com/watch?v=JpeHIbzmGP4

https://redis.io/topics/streams-intro

https://redislabs.com/blog/use-redis-streams-apps/

https://redislabs.com/blog/getting-started-with-redis-streams-and-java/

