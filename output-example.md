# DBdog
## posts
|Field|Type|Null|Key|Default|Extra|
|---|---|---|---|---|---|
|id|int(10) unsigned|NO|PRI||auto_increment|
|user_id|int(10) unsigned|NO|MUL|||
|title|varchar(255)|NO||||
|body|text|YES||||
|published_at|timestamp|YES||||
|deleted_at|timestamp|YES||||
|created_at|timestamp|NO||CURRENT_TIMESTAMP|DEFAULT_GENERATED|
|updated_at|timestamp|NO||CURRENT_TIMESTAMP|DEFAULT_GENERATED on update CURRENT_TIMESTAMP|

## users
|Field|Type|Null|Key|Default|Extra|
|---|---|---|---|---|---|
|id|int(10) unsigned|NO|PRI||auto_increment|
|name|varchar(255)|NO||||
|email|varchar(255)|NO|UNI|||
|email_verified_at|timestamp|YES||||
|password|varchar(255)|NO||||
|remember_token|varchar(100)|YES||||
|status_id|int(11)|YES||0||
|created_at|timestamp|NO||CURRENT_TIMESTAMP|DEFAULT_GENERATED|
|updated_at|timestamp|NO||CURRENT_TIMESTAMP|DEFAULT_GENERATED on update CURRENT_TIMESTAMP|

