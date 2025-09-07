#!/bin/sh
set -euo pipefail

MIGRATIONS_DIR=${MIGRATIONS_DIR:-/migrations}
RETRIES=${DB_MIGRATION_RETRIES:-30}
SLEEP=${DB_MIGRATION_SLEEP:-2}
APPLY=${MIGRATIONS_APPLY:-1}

run_migrations() {
  echo "[migrate] Applying migrations from $MIGRATIONS_DIR";
  i=0
  until atlas migrate apply --dir "file://$MIGRATIONS_DIR" --url "$DATABASE_URL"; do
    i=$((i+1))
    if [ "$i" -ge "$RETRIES" ]; then
      echo "[migrate] Failed after $RETRIES attempts" >&2
      return 1
    fi
    echo "[migrate] Retry $i/$RETRIES in ${SLEEP}s..."
    sleep "$SLEEP"
  done
  echo "[migrate] Done"
}

if [ "$APPLY" = "1" ]; then
  if [ -z "${DATABASE_URL:-}" ]; then
    echo "[migrate] DATABASE_URL not set; skipping migrations" >&2
  else
    run_migrations
  fi
else
  echo "[migrate] Skipped (MIGRATIONS_APPLY=$APPLY)"
fi

echo "[app] Starting application"
exec /app

