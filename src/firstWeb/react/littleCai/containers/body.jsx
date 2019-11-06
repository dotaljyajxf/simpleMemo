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
          <div className="carousel-one carousel-img"></div>
          <div className="carousel-two carousel-img"></div>
          <div className="carousel-three carousel-img"></div>
          <div className="carousel-four carousel-img"></div>
          <div className="carousel-five carousel-img"></div>
        </Carousel>
      </div>
    );
  }
}

export default Body;
