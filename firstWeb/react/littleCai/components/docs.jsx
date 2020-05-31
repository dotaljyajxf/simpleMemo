import React, { Component } from "react";
import { Carousel } from "antd";

const zh = {
  hello: "Hello",
  login: "登录",
  message: "请输入信息",
  docName1: "doc1",
  docName2: "doc2",
};

class Docs extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  render() {
    return (
      <div>
        <Carousel autoplay>
          <div className="carousel-img">{zh.docName1}</div>
          <div className="carousel-img">{zh.docName2}</div>
        </Carousel>
      </div>
    );
  }
}

export default Docs;
