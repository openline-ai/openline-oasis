#!/bin/bash
FILES="standard-create.sql permissions-create.sql carriers.sql"
if [ "x$1" == "xlocal-kube" ]; then
  pod=$(kubectl get pods -n openline-development|grep oasis-postgres| cut -f1 -d ' ')
  echo $FILES |xargs cat|kubectl exec -it $pod -- psql $SQL_USER $SQL_DATABASE
else
  echo $FILES |xargs cat| PGPASSWORD=$SQL_PASSWORD  psql -h $SQL_HOST $SQL_USER $SQL_DATABASE
fi
