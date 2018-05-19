LINK=http://7ep7acrkunzdcw3l.onion
OUTPUT=output.json
CONCURRENCE=10

run:
	hack/run.sh $(LINK) $(OUTPUT) $(CONCURRENCE)

check-integration:
	hack/check-integration.sh

check:
	hack/check.sh