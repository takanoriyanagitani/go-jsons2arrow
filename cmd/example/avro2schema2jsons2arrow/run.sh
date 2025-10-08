#!/bin/sh

geninput(){

	jq -c -n '{
		timestamp:"2025-10-06T00:19:40.012345Z", 
		unixtime: 1234567890,
		severity:"INFO",
		status:200,
		value:299792458,
		amount:42.195,
		height:3.776,
		method:"GET",
		body:"apt update done",
		id:"cafef00d-dead-beaf-face-864299792458",
		active:true,
	}'

	jq -c -n '{
		timestamp:"2025-10-07T00:19:40.012345Z", 
		unixtime: 1234567891,
		severity:"WARN",
		status:200,
		value:16777216,
		amount:42.195,
		height: 0.599,
		method:"GET",
		author:"jd",
		body:"apt update failure",
		id:"cafef00d-dead-beaf-face-864299792458",
		active:false,
	}'

}

genschema(){
	jq -n '{
		type:"record",
		name:"log",
		fields:[
			{name:"timestamp",type:"string"},
			{name:"unixtime", type: {
				type:"long",
				logicalType:"timestamp-micros",
			}},
			{name:"severity",type:{
				type:"enum",
				name:"Severity",
				symbols:[
					"TRACE",
					"DEBUG",
					"INFO",
					"WARN",
					"ERROR",
					"FATAL"
				]
			}},
			{name:"status",type:"int"},
			{name:"value",type:"long"},
			{name:"amount",type:"float"},
			{name:"height",type:"double"},
			{name:"method",type:"string"},
			{name:"active",type:"boolean"},
			{name:"author",type:["null","string"]},
			{name:"id",type:{
				type:"string",
				logicalType:"uuid",
			}},
			{name:"body",type:"string"}
		],
	}'
}

export ENV_AVRO_SCHEMA_NAME=./sample.avsc

test -f "${ENV_AVRO_SCHEMA_NAME}" || genschema > "${ENV_AVRO_SCHEMA_NAME}"

geninput |
	./avro2schema2jsons2arrow
