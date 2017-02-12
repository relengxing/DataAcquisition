package com.relengxing;

import com.relengxing.entity.MeterInfo;
import com.relengxing.mapper.Date1c20Mapper;
import com.relengxing.mapper.MeterInfoMapper;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

@RunWith(SpringRunner.class)
@SpringBootTest
public class WebApplicationTests {

	@Autowired
	Date1c20Mapper date1c20Mapper;

	@Autowired
	MeterInfoMapper meterInfoMapper;
	@Test
	public void contextLoads() {
//		System.out.println(date1c20Mapper.findAllAddress());
		System.out.println(meterInfoMapper.findByAddress(
				"1a1b236a4e17"));
	}

}
