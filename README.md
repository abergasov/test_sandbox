## Overview

local run in docker
```shell
make run
```

```bash
POST /api/messages?per_page=10&page=1
GET /api/all_messages
GET /api/message/:id
DELETE /api/message/:id"

# update message
curl -X POST http://127.0.0.1:8000/api/message/316253 -header "Content-Type: application/json" --data '{"new_text":"xyz","is_bot":true}'
```