import React, { Component } from "react";
import '../style/body.css';

const zh = {
    design: 'Made By 嘿哈小组'
};

class FooterContent extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  render() {
    return (
      <div className="footer-content">
        <span className="made-by-info">{zh.design}</span>
      </div>
    );
  }
}

export default FooterContent;
