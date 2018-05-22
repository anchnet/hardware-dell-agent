#!/usr/bin/env bash
HOST=192.168.106.105
USER=root
PWD=root
ipmitool -H $HOST -I lanplus -U $USER -P $PWD -v sdr list|while read line
do
if [[ $line =~ 'Sensor ID' ]];then
sensor_id=$(echo $line|awk -F':' '{print $2}')
fi
if [[ $line =~ 'Entity ID' ]];then
entity_id=$(echo $line|awk -F':' '{print $2}')
fi
if [[ $line =~ 'Sensor Type' ]];then
sensor_type=$(echo $line|awk -F':' '{print $2}')
fi
if [[ $line =~ 'Sensor Reading' ]];then
sensor_reading=$(echo $line|awk -F':' '{print $2}')
fi
if [[ $line =~ 'Status' ]];then
sensor_status=$(echo $line|awk -F':' '{print $2}')
fi
if [[ $line =~ 'Lower critical' ]];then
lower_crit=$(echo $line|awk -F':' '{print $2}')
fi
if [[ $line =~ 'Lower non-critical' ]];then
lower_non_crit=$(echo $line|awk -F':' '{print $2}')
fi
if [[ $line =~ 'Upper critical' ]];then
upper_crit=$(echo $line|awk -F':' '{print $2}')
fi
if [[ $line =~ 'Upper non-critical' ]];then
upper_non_crit=$(echo $line|awk -F':' '{print $2}')
fi
if [[ $line = '' ]];then
    echo $entity_id\|$sensor_id\|$sensor_type\|$sensor_reading\|$lower_crit\|$lower_non_crit\|$upper_crit\|$upper_non_crit\|$sensor_status
    sensor_status=''
    lower_crit=''
    lower_non_crit=''
    upper_crit=''
    upper_non_crit=''
fi
#if [[ -n $entity_id ]] && [[ -n $sensor_id ]] && [[ -n $sensor_type ]] && [[ -n $sensor_reading ]];then
#    echo $entity_id\|$sensor_id\|$sensor_type\|$sensor_reading
#fi
done
