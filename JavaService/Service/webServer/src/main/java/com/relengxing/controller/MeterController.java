package com.relengxing.controller;

import com.alibaba.fastjson.JSON;
import com.relengxing.entity.Date1C20;
import com.relengxing.entity.MeterInfo;
import com.relengxing.mapper.Date1c20Mapper;
import com.relengxing.mapper.MeterInfoMapper;
import org.apache.ibatis.annotations.Param;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * Created by relengxing on 2017/2/4.
 */
@RestController
public class MeterController {

    @Autowired
    Date1c20Mapper date1c20Mapper;

    @Autowired
    MeterInfoMapper meterInfoMapper;
    /**
     * 获取所有存在数据的表记地址
     * */
    @RequestMapping("/meter/1c20/address")
    public String getAllMeter(){
        return date1c20Mapper.findAllAddress().toString();
    }

    /**
     * 获取所有存在数据的表信息
     * */
    @RequestMapping("/meter/map")
    public String getAllLocation(){
        List<String> address = date1c20Mapper.findAllAddress();
        Map<String,String> map = new HashMap<>();
        for (String addr:address) {
            String locat = meterInfoMapper.findLocationByAddress(addr);
            map.put(addr,locat);
        }
        return JSON.toJSONString(map);
    }
    /**
     * 获取所有表信息（不管是否存在数据）
     * */
    @RequestMapping(value = "/meter",method = RequestMethod.GET)
    public String getAllMeterInfo(){
        List<MeterInfo> meterInfos = meterInfoMapper.findAll();
        return JSON.toJSONString(meterInfos);
    }

    @RequestMapping(value = "/meter/{address}",method = RequestMethod.GET)
    public String getMeterInfo(@PathVariable("address")String address){
        MeterInfo meterInfo = meterInfoMapper.findByAddress(address);
        String json = JSON.toJSONString(meterInfo);
        return json;
    }

    /**
     * 获取某个地址的所有数据
     * */
    @RequestMapping(value = "/meter/{address}/1c20",method = RequestMethod.GET)
    public String getMeterData(@PathVariable("address")String address){
        List<Date1C20> list = date1c20Mapper.findByAddress(address);
        String json = JSON.toJSONString(list);
        return json;
    }

}
