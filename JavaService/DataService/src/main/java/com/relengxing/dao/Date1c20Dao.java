package com.relengxing.dao;

import com.relengxing.bean.Date1c20;
import com.relengxing.mapper.Date1c20Mapper;
import org.apache.ibatis.io.Resources;
import org.apache.ibatis.session.SqlSession;
import org.apache.ibatis.session.SqlSessionFactory;
import org.apache.ibatis.session.SqlSessionFactoryBuilder;

import java.io.IOException;
import java.io.InputStream;
import java.util.List;

/**
 * Created by relengxing on 2016/12/28.
 */
public class Date1c20Dao {
    private static String resource;
    private static SqlSessionFactory factory;
    static {
        resource = "mybatis-config.xml";
        InputStream inputStream = null;
        try {
            inputStream = Resources.getResourceAsStream(resource);
        } catch (IOException e) {
            e.printStackTrace();
        }
        factory = new SqlSessionFactoryBuilder().build(inputStream);
    }
    /**
     * 增
     * */
    public static void add(Date1c20 date1c20){
        SqlSession sqlSession = factory.openSession();
        Date1c20Mapper mapper = sqlSession.getMapper(Date1c20Mapper.class);
        mapper.insertSelective(date1c20);
        sqlSession.commit();
        sqlSession.close();
    }
    /***
     * 删
     */
    public static void deleteById(int id){
        SqlSession sqlSession = factory.openSession();
        Date1c20Mapper mapper = sqlSession.getMapper(Date1c20Mapper.class);
        mapper.deleteByPrimaryKey(id);
        sqlSession.close();
    }
    /**
     * 改
     * */
    public static void updateById(Date1c20 record){
        SqlSession sqlSession = factory.openSession();
        Date1c20Mapper mapper = sqlSession.getMapper(Date1c20Mapper.class);
        mapper.updateByPrimaryKey(record);
        sqlSession.close();
    }
    /**
     * 查
     * **/
    public static Date1c20 getById(int id){
        SqlSession sqlSession = factory.openSession();
        Date1c20Mapper mapper = sqlSession.getMapper(Date1c20Mapper.class);
        Date1c20 date1c20 = mapper.selectByPrimaryKey(1);
        sqlSession.close();
        return date1c20;
    }

    public static List<Date1c20> getAll(){
        SqlSession sqlSession = factory.openSession();
        Date1c20Mapper mapper = sqlSession.getMapper(Date1c20Mapper.class);
        List<Date1c20> date1c20 = mapper.selectAllDate1c20();
        sqlSession.close();
        return date1c20;
    }
}
