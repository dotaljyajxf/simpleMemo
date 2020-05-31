import React, { Component } from "react";
import { Calendar } from 'antd';

class CalendarComponent extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
    this.onPanelChange = this.onPanelChange.bind(this);
  }

  onPanelChange(value, mode) {
    console.log(value, mode);
  }

  render() {
    return (
        <Calendar onPanelChange={this.onPanelChange} />
    );
  }
}

export default CalendarComponent;
