import React from "react";
import './index.less';
import 'echarts/lib/echarts';
import ReactEcharts from 'echarts-for-react';

const ContrastMonitor = (props) => {
    const { list4 } = props;

    let base = +new Date(1988, 9, 3);
    let oneDay = 24 * 3600 * 1000;
    let data = [[base, Math.random() * 300]];
    for (let i = 1; i < 20000; i++) {
        let now = new Date((base += oneDay));
        data.push([+now, Math.round((Math.random() - 0.5) * 20 + data[i - 1][1])]);
    }

    let option = {
        tooltip: {
          trigger: 'axis',
          position: function (pt) {
            return [pt[0], '10%'];
          }
        },
        title: {
          left: 'center',
          text: 'Large Ara Chart'
        },
        xAxis: {
          type: 'time',
          boundaryGap: false
        },
        yAxis: {
          type: 'value',
          boundaryGap: [0, '100%']
        },
        series: [
          {
            name: 'Fake Data',
            type: 'line',
            smooth: true,
            symbol: 'none',
            areaStyle: {},
            data: data
          }
        ]
      };

    return (
        <div className="contrast-monitor">
            <div className="contrast-monitor-list">
                <ReactEcharts className='echarts' option={option} />
                <ReactEcharts className='echarts' option={option} />
                <ReactEcharts className='echarts' option={option} />
                <ReactEcharts className='echarts' option={option} />
            </div>
        </div>
    )
};

export default ContrastMonitor;