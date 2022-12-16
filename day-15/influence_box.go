package main

type InfluenceBox struct {
	max Vec2
	min Vec2
}

func GetInfluenceBoxForSensor(sensor *Sensor) InfluenceBox {
	distance := sensor.GetDistance()
	return InfluenceBox{
		max: NewVec2(
			sensor.position.x+distance,
			sensor.position.y+distance,
		),
		min: NewVec2(
			sensor.position.x-distance,
			sensor.position.y-distance,
		),
	}
}
