package com.jbennett.demo;

import java.util.Map;

import com.google.common.collect.ImmutableMap;
import org.apache.commons.lang3.StringUtils;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@SpringBootApplication
@RestController
public class InventoryApplication {

    public static void main(String[] args) {
        SpringApplication.run(InventoryApplication.class, args);
    }

    @GetMapping("/api/stock")
    public Map<String, Object> stock() {
        return ImmutableMap.of(
                "service", StringUtils.capitalize("inventory"),
                "status", "ok");
    }
}
