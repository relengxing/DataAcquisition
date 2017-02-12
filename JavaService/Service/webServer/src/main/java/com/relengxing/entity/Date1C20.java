package com.relengxing.entity;

import java.util.Date;

/**
 * Created by relengxing on 2017/2/3.
 */
public class Date1C20 {

    private int id;

    private String address;

    private Date time;

    private int temerature;

    private int humidity;

    private int electricity;

    private String location;

    public String getLocation() {
        return location;
    }

    public void setLocation(String location) {
        this.location = location;
    }

    public Date1C20() {
    }

    public Date1C20(int id, String address, Date time, int temerature, int humidity, int electricity) {
        this.id = id;
        this.address = address;
        this.time = time;
        this.temerature = temerature;
        this.humidity = humidity;
        this.electricity = electricity;
    }

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getAddress() {
        return address;
    }

    public void setAddress(String address) {
        this.address = address;
    }

    public Date getTime() {
        return time;
    }

    public void setTime(Date time) {
        this.time = time;
    }

    public int getTemerature() {
        return temerature;
    }

    public void setTemerature(int temerature) {
        this.temerature = temerature;
    }

    public int getHumidity() {
        return humidity;
    }

    public void setHumidity(int humidity) {
        this.humidity = humidity;
    }

    public int getElectricity() {
        return electricity;
    }

    public void setElectricity(int electricity) {
        this.electricity = electricity;
    }

    @Override
    public String toString() {
        return "Date1C20{" +
                "id=" + id +
                ", address='" + address + '\'' +
                ", time=" + time +
                ", temerature=" + temerature +
                ", humidity=" + humidity +
                ", electricity=" + electricity +
                ", location='" + location + '\'' +
                '}';
    }
}
