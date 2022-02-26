# Deploy process
1. Write secrets to secrets file with `echo -n`
- `echo -n MONGO_INITDB_ROOT_USERNAME > bap_db/secrets/.mongo_root`
- `echo -n MONGO_INITDB_ROOT_PASSWORD > bap_db/secrets/.mongo_root_password`
- `echo -n SLACK_SIGNING_SECRET > bap_back/secrets/.slack_signing_secret`
- `echo -n SLACK_BOT_USER_OAUTH_TOKEN > bap_back/secrets/.slack_bot_user_oauth_token`
- `echo -n GOOGLE_APPLICATION_CREDENTIALS > bap_back/secrets/.drive_api_service_account`

2. Create mount directory for mongo docker
- `mkdir bap_db/db`

3. Optional: Write data to initialize database
- `bap_db/init/blog.json`
- `bap_db/init/profile.json`

4. Deploy with docker-compose
- `docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d`
