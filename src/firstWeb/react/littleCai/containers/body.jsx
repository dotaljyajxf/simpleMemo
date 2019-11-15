import React, { Component } from "react";
import { Carousel } from "antd";
import '../style/body.css';

const zh = {
  hello: "Hello",
  login: "登录",
  message: "请输入信息"
};

class Body extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  render() {
    return (
      <div className="body-content">
        <Carousel autoplay>
          <div className="carousel-img carousel-one"></div>
          <div className="carousel-img carousel-two"></div>
        </Carousel>
      </div>
    );
  }
}

export default Body;
