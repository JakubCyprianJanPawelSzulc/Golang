import { useEffect, useState } from 'react';
import { booksService } from '../../../services/books.service';
import './StatisticsActions.scss';
import { Chart } from 'primereact/chart';

export function StatisticsActions() {
  const [mostBooks, setMostBooks] = useState(null);
  const [oldestBooks, setOldestBooks] = useState(null);
  const [mostPopular, setMostPopularBooks] = useState(null);
  const [chartData, setChartData] = useState();
  const [chartRadarData, setChartRadarData] = useState();
  const [barData, setBarData] = useState();

  useEffect(() => {
    booksService.getTenMostBooks().then((e) => {
      setMostBooks(e);
    });

    booksService.getTenOldestBooks().then((e) => {
      console.log(e);
      setOldestBooks(e);
    });

    booksService.getFiveMostPopularBooks().then((e) => {
      setMostPopularBooks(e);
    });
  }, []);

  useEffect(() => {
    if (mostBooks) {
      const qqt = mostBooks.map((e) => e.quantity);
      const titles = mostBooks.map((e) => e.title);

      setChartData({
        labels: titles,
        datasets: [
          {
            data: qqt,
            backgroundColor: [
              '#FF6384',
              '#36A2EB',
              '#FFCE56',
              '#673AB7',
              '#795548',
              '#E040FB',
              '#00E676',
              '#78909C',
              '#FF5252',
              '#FF7043'
            ],
            hoverBackgroundColor: [
              '#FF6384',
              '#36A2EB',
              '#FFCE56',
              '#673AB7',
              '#795548',
              '#E040FB',
              '#00E676',
              '#78909C',
              '#FF5252',
              '#FF7043'
            ]
          }
        ]
      });
    }
  }, [mostBooks]);

  useEffect(() => {
    if (mostPopular) {
      console.log('PPR', mostPopular);
      const qqty = mostPopular.map((e) => e.reservationCount);
      const titles = mostPopular.map((e) => e.title);

      console.log(qqty);

      setChartRadarData({
        labels: titles,
        datasets: [
          {
            data: qqty,
            backgroundColor: [
              '#FF6384',
              '#36A2EB',
              '#FFCE56',
              '#673AB7',
              '#795548',
              '#E040FB',
              '#00E676',
              '#78909C',
              '#FF5252',
              '#FF7043'
            ],
            hoverBackgroundColor: [
              '#FF6384',
              '#36A2EB',
              '#FFCE56',
              '#673AB7',
              '#795548',
              '#E040FB',
              '#00E676',
              '#78909C',
              '#FF5252',
              '#FF7043'
            ]
          }
        ]
      });
    }
  }, [mostPopular]);

  useEffect(() => {
    if (oldestBooks) {
      const years = oldestBooks.map((e) => new Date(e.dateRelease).getFullYear());
      const titles = oldestBooks.map((e) => e.title);

      setBarData({
        labels: titles,
        datasets: [
          {
            label: 'Najstarsze książki',
            backgroundColor: '#42A5F5',
            data: years
          }
        ]
      });
    }
  }, [oldestBooks]);

  const [lightOptions] = useState({
    plugins: {
      legend: {
        labels: {
          color: '#495057'
        }
      }
    }
  });

  let basicOptions = {
    maintainAspectRatio: false,
    aspectRatio: 0.8,

    plugins: {
      legend: {
        labels: {
          color: '#495057'
        }
      }
    },
    scales: {
      x: {
        ticks: {
          color: '#495057'
        },
        grid: {
          color: '#ebedef'
        }
      },
      y: {
        ticks: {
          stepSize: 1000,
          color: '#495057'
        },
        grid: {
          color: '#ebedef'
        }
      }
    }
  };

  return (
    <div className="charts">
      <div className="top-charts">
        <div className="label-b">Najwięcej książek</div>

        <div className="label-b">Najbardziej popularne</div>
      </div>
      <div className="top-charts">
        {mostBooks && <Chart type="doughnut" data={chartData} style={{ position: 'relative', width: '40%' }} />}

        <Chart
          type="polarArea"
          data={chartRadarData}
          options={lightOptions}
          style={{ position: 'relative', width: '40%' }}
        />
      </div>
      <Chart type="bar" data={barData} options={basicOptions} />
    </div>
  );
}
