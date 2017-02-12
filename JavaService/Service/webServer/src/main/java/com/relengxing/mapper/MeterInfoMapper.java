package com.relengxing.mapper;

import com.relengxing.entity.MeterInfo;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;
import org.apache.ibatis.annotations.Select;
import org.springframework.stereotype.Component;

import java.util.List;

/**
 * Created by relengxing on 2017/2/4.
 */
@Mapper
@Component
public interface MeterInfoMapper {

    @Select("SELECT * FROM meterinfo")
    List<MeterInfo> findAll();

    @Select("SELECT * FROM meterinfo WHERE address=#{address}")
    MeterInfo findByAddress(String address);

    @Select("SELECT location FROM meterinfo where address=#{address}")
    String findLocationByAddress(@Param("address")String address);
}
