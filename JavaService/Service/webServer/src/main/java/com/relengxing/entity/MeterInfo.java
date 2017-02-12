package com.relengxing.entity;

/**
 * Created by relengxing on 2017/2/4.
 */
public class MeterInfo {

    private int id;

    private String location;

    private String address;


    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getLocation() {
        return location;
    }

    public void setLocation(String location) {
        this.location = location;
    }

    public String getAddress() {
        return address;
    }

    public void setAddress(String address) {
        this.address = address;
    }

    @Override
    public String toString() {
        return "MeterInfo{" +
                "id=" + id +
                ", location='" + location + '\'' +
                ", address='" + address + '\'' +
                '}';
    }
}
