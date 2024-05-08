filepath=vars/.env

docker compose --env-file="$filepath" up -d

echo "---config start--"
docker compose --env-file="$filepath" config
echo "---Config end --"