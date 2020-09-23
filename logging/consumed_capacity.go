package logging

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"go.uber.org/zap"
)

func DynamoDBConsumedCapacityFields(ccs ...*dynamodb.ConsumedCapacity) []zap.Field {
	fields := []zap.Field{}
	if ccs == nil || len(ccs) == 0 {
		return fields
	}
	for i, cc := range ccs {
		prefix := fmt.Sprintf("consumedCapacity.%02d.", i)
		fields = append(fields,
			zap.Stringp(prefix+"tableName", cc.TableName),
			zap.Float64p(prefix+"totalCapacityUnits", cc.CapacityUnits),
			zap.Float64p(prefix+"readCapacityUnits", cc.ReadCapacityUnits),
			zap.Float64p(prefix+"writeCapacityUnits", cc.WriteCapacityUnits),
		)
		if cc.Table != nil {
			fields = append(
				fields,
				zap.Float64p(prefix+"table.totalCapacityUnits", cc.Table.CapacityUnits),
				zap.Float64p(prefix+"table.readCapacityUnits", cc.Table.ReadCapacityUnits),
				zap.Float64p(prefix+"table.writeCapacityUnits", cc.Table.WriteCapacityUnits),
			)
		}
		if cc.GlobalSecondaryIndexes != nil {
			for name, capacity := range cc.GlobalSecondaryIndexes {
				localPrefix := prefix + "gsi." + name
				fields = append(
					fields,
					zap.Float64p(localPrefix+".totalCapacityUnits", capacity.CapacityUnits),
					zap.Float64p(localPrefix+".readCapacityUnits", capacity.ReadCapacityUnits),
					zap.Float64p(localPrefix+".writeCapacityUnits", capacity.WriteCapacityUnits),
				)
			}
		}
		if cc.LocalSecondaryIndexes != nil {
			for name, capacity := range cc.LocalSecondaryIndexes {
				localPrefix := prefix + "lsi." + name
				fields = append(
					fields,
					zap.Float64p(localPrefix+".totalCapacityUnits", capacity.CapacityUnits),
					zap.Float64p(localPrefix+".readCapacityUnits", capacity.ReadCapacityUnits),
					zap.Float64p(localPrefix+".writeCapacityUnits", capacity.WriteCapacityUnits),
				)
			}
		}
	}
	return fields
}
