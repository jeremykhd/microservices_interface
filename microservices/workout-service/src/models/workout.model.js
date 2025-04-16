const mongoose = require('mongoose');

const exerciseSchema = new mongoose.Schema({
  name: {
    type: String,
    required: true
  },
  sets: {
    type: Number,
    required: true
  },
  reps: {
    type: Number,
    required: true
  },
  weight: {
    type: Number,
    required: false
  },
  duration: {
    type: Number, // in minutes
    required: false
  },
  notes: {
    type: String,
    required: false
  }
});

const workoutSchema = new mongoose.Schema({
  userId: {
    type: mongoose.Schema.Types.ObjectId,
    required: true,
    ref: 'User'
  },
  date: {
    type: Date,
    required: true,
    default: Date.now
  },
  name: {
    type: String,
    required: true
  },
  type: {
    type: String,
    enum: ['strength', 'cardio', 'flexibility', 'mixed'],
    required: true
  },
  duration: {
    type: Number, // in minutes
    required: true
  },
  exercises: [exerciseSchema],
  notes: {
    type: String,
    required: false
  },
  caloriesBurned: {
    type: Number,
    required: false
  },
  difficulty: {
    type: String,
    enum: ['easy', 'medium', 'hard'],
    required: true
  }
}, {
  timestamps: true
});

module.exports = mongoose.model('Workout', workoutSchema); 