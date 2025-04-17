# RedkaCli

## Adaptation list

### Strings

| Redis Command    | redkacli.Client    | redka.DB            | Description                                       |
|-------------|--------------------|---------------------|--------------------------------------------|
| DECR        | client.Decr        | DB.Str().Incr       | Decrements the integer value of a key by one.|
| DECRBY      | client.DecrBy      | DB.Str().Incr       | Decrements a number from the integer value of a key.|
| GET         | client.Get         | DB.Str().Get        | Returns the value of a key.                |
| GETSET      | client.GetSet      | DB.Str().SetWith    | Sets the key to a new value and returns the prev value.|
| INCR        | client.Incr        | DB.Str().Incr       | Increments the integer value of a key by one.|
| INCRBY      | client.IncrBy      | DB.Str().Incr       | Increments the integer value of a key by a number.|
| INCRBYFLOAT | client.IncrByFloat | DB.Str().IncrFloat  | Increments the float value of a key by a number.|
| MGET        | client.MGet        | DB.Str().GetMany    | Returns the values of one or more keys.      |
| MSET        | client.MSet        | DB.Str().SetMany    | Sets the values of one or more keys.         |
| PSETEX      | client.PSetEx      | DB.Str().SetExpires | Sets the value and expiration time (in ms) of a key.|
| SET         | client.Set         | DB.Str().Set        | Sets the value of a key.                     |
| SETEX       | client.SetEx       | DB.Str().SetExpires | Sets the value and expiration (in sec) time of a key.|
| SETNX       | client.SetNX       | DB.Str().SetWith    | Sets the value of a key when the key doesn't exist.|
| STRLEN      | client.StrLen      | DB.Str().Get        | Returns the length of a value in bytes.      |


### Hashes

| Redis Command    | redkacli.Client     | redka.DB                             | Description                                       |
|-------------|---------------------|--------------------------------------|--------------------------------------------|
| HDEL        | client.HDel         | DB.Hash().Delete                     | Deletes one or more fields and their values.|
| HEXISTS       | client.HExists      | DB.Hash().Exists                     | Determines whether a field exists.         |
| HGET        | client.HGet         | DB.Hash().Get                        | Returns the value of a field.              |
| HGETALL       | client.HGetAll      | DB.Hash().Items                      | Returns all fields and values.           |
| HINCRBY       | client.HincrBy      | DB.Hash().Incr                       | Increments the integer value of a field.   |
| HINCRBYFLOAT  | client.HIncrByFloat | DB.Hash().IncrFloat                  | Increments the float value of a field.     |
| HKEYS         | client.HKeys        | DB.Hash().Keys => DB.Hash().Fields   | Returns all fields.                      |
| HLEN          | client.HLen         | DB.Hash().Len                        | Returns the number of fields.              |
| HMGET         | client.HMGet        | DB.Hash().GetMany                    | Returns the values of multiple fields.     |
| HMSET         | client.HMSet        | DB.Hash().SetMany                    | Sets the values of multiple fields.        |
| HSCAN         | client.HScan        | DB.Hash().Scanner                    | Iterates over fields and values.         |
| HSET          | client.HSet         | DB.Hash().SetMany                    | Sets the values of one or more fields.     |
| HSETNX        | client.HSetNX       | DB.Hash().SetNotExists               | Sets the value of a field when it doesn't exist.|
| HVALS         | client.HVals             | DB.Hash().Exists => DB.Hash().Values | Returns all values.                        |


### Key management

| Redis Command    | redkacli.Client | redka.DB            | Description                                       |
|-------------|-----------------|---------------------|--------------------------------------------|
| DBSIZE        |                 | DB.Key().Len        | Returns the total number of keys.          |
| DEL           |                 | DB.Key().Delete     | Deletes one or more keys.                  |
| EXISTS        |                 | DB.Key().Count      | Determines whether one or more keys exist. |
| EXPIRE        |                 | DB.Key().Expire     | Sets the expiration time of a key (in seconds).|
| EXPIREAT      |                 | DB.Key().ExpireAt   | Sets the expiration time of a key to a Unix timestamp.|
| FLUSHALL      |                 | DB.Key().DeleteAll  | Deletes all keys from the database.        |
| FLUSHDB       |                 | DB.Key().DeleteAll  | Deletes all keys from the database.        |
| KEYS          |                 | DB.Key().Keys       | Returns all key names that match a pattern.|
| PERSIST       |                 | DB.Key().Persist    | Removes the expiration time of a key.      |
| PEXPIRE       |                 | DB.Key().Expire     | Sets the expiration time of a key in ms.   |
| PEXPIREAT     |                 | DB.Key().ExpireAt   | Sets the expiration time of a key to a Unix ms timestamp.|
| RANDOMKEY     |                 | DB.Key().Random     | Returns a random key name from the database.|
| RENAME        |                 | DB.Key().Rename     | Renames a key and overwrites the destination.|
| RENAMENX      |                 | DB.Key().RenameNotExists | Renames a key only when the target key name doesn't exist.|
| SCAN          |                 | DB.Key().Scanner    | Iterates over the key names in the database.|
| TTL           |                 | DB.Key().Get        | Returns the expiration time in seconds of a key.|
| TYPE          |                 | DB.Key().Get        | Returns the type of value stored at a key. |


### Lists

| Redis Command    | redkacli.Client | redka.DB            | Description                                       |
|-------------|-----------------|---------------------|--------------------------------------------|
| LINDEX        |                 | DB.List().Get       | Returns an element by its index.           |
| LINSERT       |                 | DB.List().Insert*   | Inserts an element before or after another element.|
| LLEN          |                 | DB.List().Len       | Returns the length of a list.              |
| LPOP          |                 | DB.List().PopFront  | Returns the first element after removing it.|
| LPUSH         |                 | DB.List().PushFront | Prepends an element to a list.             |
| LRANGE        |                 | DB.List().Range     | Returns a range of elements.               |
| LREM          |                 | DB.List().Delete*   | Removes elements from a list.              |
| LSET          |                 | DB.List().Set       | Sets the value of an element by its index.   |
| LTRIM         |                 | DB.List().Trim      | Removes elements from both ends of a list. |
| RPOP          |                 | DB.List().PopBack   | Returns the last element after removing it.|
| RPOPLPUSH     |                 | DB.List().PopBackPushFront | Removes the last element and pushes it to another list.|
| RPUSH         |                 | DB.List().PushBack  | Appends an element to a list.              |


### Server/connection management

| Redis Command | redkacli.Client | redka.DB            | Description                                       |
|------------|-----------------|---------------------|--------------------------------------------|
| ECHO       |                 |                     | Returns the given string.                  |
| LOLWUT     |                 |                     | Provides an answer to a yes/no question.     |
| PING       |                 |                     | Returns the server's liveliness response.    |
| SELECT     |                 |                     | Changes the selected database (no-op).     |


### Sets

| Redis Command | redkacli.Client | redka.DB            | Description                                       |
|------------|-----------------|---------------------|--------------------------------------------|
| SADD       |                 | DB.Set().Add        | Adds one or more members to a set.         |
| SCARD      |                 | DB.Set().Len        | Returns the number of members in a set.    |
| SDIFF      |                 | DB.Set().Diff       | Returns the difference of multiple sets.   |
| SDIFFSTORE |                 | DB.Set().DiffStore  | Stores the difference of multiple sets.    |
| SINTER     |                 | DB.Set().Inter      | Returns the intersection of multiple sets. |
| SINTERSTORE|                 | DB.Set().InterStore | Stores the intersection of multiple sets.  |
| SISMEMBER  |                 | DB.Set().Exists     | Determines whether a member belongs to a set.|
| SMEMBERS   |                 | DB.Set().Items      | Returns all members of a set.              |
| SMOVE      |                 | DB.Set().Move       | Moves a member from one set to another.    |
| SPOP       |                 | DB.Set().Pop        | Returns a random member after removing it. |
| SRANDMEMBER|                 | DB.Set().Random     | Returns a random member from a set.        |
| SREM       |                 | DB.Set().Delete     | Removes one or more members from a set.    |
| SSCAN      |                 | DB.Set().Scanner    | Iterates over members of a set.            |
| SUNION     |                 | DB.Set().Union      | Returns the union of multiple sets.        |
| SUNIONSTORE|                 | DB.Set().UnionStore | Stores the union of multiple sets.         |


### Zsets

| Redis Command      | redkacli.Client | redka.DB            | Description                                       |
|-----------------|-----------------|---------------------|--------------------------------------------|
| ZADD            |                 | DB.ZSet().AddMany   | Adds or updates one or more members of a set.|
| ZCARD           |                 | DB.ZSet().Len       | Returns the number of members in a set.      |
| ZCOUNT          |                 | DB.ZSet().Count     | Returns the number of members of a set within a range of scores.|
| ZINCRBY         |                 | DB.ZSet().Incr      | Increments the score of a member in a set.   |
| ZINTER          |                 | DB.ZSet().InterWith | Returns the intersection of multiple sets.   |
| ZINTERSTORE     |                 | DB.ZSet().InterWith | Stores the intersection of multiple sets in a key.|
| ZRANGE          |                 | DB.ZSet().RangeWith | Returns members of a set within a range of indexes.|
| ZRANGEBYSCORE   |                 | DB.ZSet().RangeWith | Returns members of a set within a range of scores.|
| ZRANK           |                 | DB.ZSet().GetRank   | Returns the index of a member in a set ordered by ascending scores.|
| ZREM            |                 | DB.ZSet().Delete    | Removes one or more members from a set.    |
| ZREMRANGEBYRANK |                 | DB.ZSet().DeleteWith| Removes members of a set within a range of indexes.|
| ZREMRANGEBYSCORE|                 | DB.ZSet().DeleteWith| Removes members of a set within a range of scores.|
| ZREVRANGE       |                 | DB.ZSet().RangeWith | Returns members of a set within a range of indexes in reverse order.|
| ZREVRANGEBYSCORE|                 | DB.ZSet().RangeWith | Returns members of a set within a range of scores in reverse order.|
| ZREVRANK        |                 | DB.ZSet().GetRankRev| Returns the index of a member in a set ordered by descending scores.|
| ZSCAN           |                 | DB.ZSet().Scan      | Iterates over members and scores of a set. |
| ZSCORE          |                 | DB.ZSet().GetScore  | Returns the score of a member in a set.    |
| ZUNION          |                 | DB.ZSet().UnionWith | Returns the union of multiple sets.        |
| ZUNIONSTORE     |                 | DB.ZSet().UnionWith | Stores the union of multiple sets in a key.  |


### Transactions

| Redis Command | redkacli.Client | redka.DB            | Description                                       |
|------------|-----------------|---------------------|--------------------------------------------|
| DISCARD    |                 | DB.View / DB.Update | Discards a transaction.                    |
| EXEC       |                 | DB.View / DB.Update | Executes all commands in a transaction.    |
| MULTI      |                 | DB.View / DB.Update | Starts a transaction.                      |

