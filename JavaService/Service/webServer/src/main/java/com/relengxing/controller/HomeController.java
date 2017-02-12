package com.relengxing.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;

/**
 * Created by relengxing on 2017/2/4.
 */

@Controller
public class HomeController {

    @RequestMapping("/")
    public String getIndex(){
        return "index";
    }

    @RequestMapping("/location")
    public String getLocation(){
        return "location";
    }
}
