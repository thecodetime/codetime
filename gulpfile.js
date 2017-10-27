var gulp = require('gulp');
var concat = require('gulp-concat');
var sass = require('gulp-sass');
var sourcemaps = require('gulp-sourcemaps');
var autoprefixer = require('gulp-autoprefixer');
var pump = require('pump');
var watch = require('gulp-watch');
var uglify = require('gulp-uglify');

gulp.task('js', (cb) => {
  pump([
    gulp.src([
      './assets/js/src/vendor/*.js',
      './assets/js/src/*.js'
    ]),
    sourcemaps.init(),
    concat('main.js'),
    uglify(),
    sourcemaps.write('.'),
    gulp.dest('./assets/js/dist')
  ], cb);
});

gulp.task('js:watch', () => {
  watch('./assets/js/src/**/*.js', () => {
    gulp.start('js');
  });
});

gulp.task('sass', (cb) => {
  pump([
    gulp.src('./assets/styles/scss/main.scss'),
    sourcemaps.init(),
    sass({
      outputStyle: 'compressed'
    }).on('error', sass.logError),
    autoprefixer(),
    sourcemaps.write('.'),
    gulp.dest('./assets/styles/dist')
  ], cb);
});

gulp.task('sass:watch', () => {
  watch('./assets/styles/scss/**/*.scss', () => {
    gulp.start('sass');
  });
});

gulp.task('default', ['js', 'js:watch', 'sass', 'sass:watch']);
gulp.task('build', ['js', 'sass']);
