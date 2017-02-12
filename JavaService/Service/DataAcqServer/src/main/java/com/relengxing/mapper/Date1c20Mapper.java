package com.relengxing.mapper;

import com.relengxing.entity.Date1C20;
import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;
import org.apache.ibatis.annotations.Select;
import org.springframework.stereotype.Component;

/**
 * Created by relengxing on 2017/2/3.
 */
@Mapper
@Component
public interface Date1c20Mapper {

    @Select("SELECT * FROM date1c20 WHERE id=#{id}")
    Date1C20 findByid(@Param("id")int id);

    @Insert("INSERT INTO date1c20(address,time,temerature,humidity,electricity) " +
            "values(#{address},#{time},#{temerature},#{humidity},#{electricity})")
    int insertData(Date1C20 date1C20);


}
