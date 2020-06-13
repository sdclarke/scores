import { Component, OnInit } from '@angular/core';
import scores from '../../assets/scores.json';

@Component({
  selector: 'app-scoreboard',
  templateUrl: './scoreboard.component.html',
  styleUrls: ['./scoreboard.component.css']
})
export class ScoreboardComponent implements OnInit {

  medianScores = [];
  maxScores = [];
  meanScores = [];

  constructor() {
    

    // for (let key in this.meanScores) {
    //   console.log(key + this.meanScores[key]);
    // }
  }

  ngOnInit() {
    this.getData();
    this.meanScores.forEach(element => {
      console.log(element + ' ' + this.meanScores[element]);
    });
  }

  getData() {
    scores.median.forEach(element => {
      this.medianScores.push({
        key:   element.name,
        value: element.score
      });
    });

    scores.maximum.forEach(element => {
      this.maxScores.push({
        key:   element.name,
        value: element.score
      });
    });

    scores.mean.forEach(element => {
      this.meanScores.push({
        key:   element.name,
        value: element.score
      });
    });
  }

}
