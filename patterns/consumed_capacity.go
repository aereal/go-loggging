package patterns

import (
	"strings"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"go.uber.org/zap"
)

func DynamoDBConsumedCapacityFields(prefix string, cc *dynamodb.ConsumedCapacity) []zap.Field {
	if !strings.HasSuffix(prefix, ".") {
		prefix = prefix + "."
	}
	if cc == nil {
		return []zap.Field{}
	}
	fields := []zap.Field{
		zap.Stringp(prefix+"tableName", cc.TableName),
		zap.Float64p(prefix+"totalCapacityUnits", cc.CapacityUnits),
		zap.Float64p(prefix+"readCapacityUnits", cc.ReadCapacityUnits),
		zap.Float64p(prefix+"writeCapacityUnits", cc.WriteCapacityUnits),
	}
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
	return fields
}
