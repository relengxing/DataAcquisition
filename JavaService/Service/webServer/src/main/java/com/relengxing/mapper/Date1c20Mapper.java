package com.relengxing.mapper;

import com.relengxing.entity.Date1C20;
import org.apache.ibatis.annotations.*;
import org.springframework.stereotype.Component;

import java.util.List;

/**
 * Created by relengxing on 2017/2/3.
 */
@Mapper
@Component
public interface Date1c20Mapper {

    @Select("SELECT * FROM date1c20 WHERE id=#{id}")
    @Results({
            @Result(id = true,property = "id",column = "id"),
            @Result(property = "location",column = "id",one = @One(select = "com.relengxing.mapper.MeterInfoMapper.findLocationByAddress")),
    })
    Date1C20 findByid(@Param("id") int id);

    @Select("SELECT distinct address FROM date1c20")
    List<String> findAllAddress();

    @Select("SELECT * FROM date1c20 WHERE address=#{address}")
    List<Date1C20> findByAddress(@Param("address")String address);

    @Insert("INSERT INTO date1c20(address,time,temerature,humidity,electricity) " +
            "values(#{address},#{time},#{temerature},#{humidity},#{electricity})")
    int insertData(Date1C20 date1C20);


}
