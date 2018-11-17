- You are setting a buffer for the channel. That is a good idea and will smooth the communication making it a bit faster. Setting a fix size would be better though. In regards to speed there is no difference between a fix size like e.g. `10` and `len(texts)`. But if `len(texts)` is huge that can become a problem as that memory needs to be allocated. If you feel a fix size might allocate too much memory most of the time, it would still be good to have an upper limit for the the channel size.