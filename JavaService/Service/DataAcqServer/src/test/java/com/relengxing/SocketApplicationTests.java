package com.relengxing;

import com.relengxing.entity.Date1C20;
import com.relengxing.mapper.Date1c20Mapper;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import java.util.Date;

@RunWith(SpringRunner.class)
@SpringBootTest
public class SocketApplicationTests {

	@Autowired
	Date1c20Mapper date1C20Mapper;
	@Test
	public void contextLoads() {
		Date1C20 date1C20 = date1C20Mapper.findByid(1);
		System.out.println(date1C20);

		Date1C20 date1C201 = new Date1C20();
		date1C201.setAddress("101010010101");
		date1C201.setTime(new Date());
		date1C201.setTemerature(365);
		date1C201.setHumidity(700);
		date1C201.setElectricity(457);
		date1C20Mapper.insertData(date1C201);
	}

}
