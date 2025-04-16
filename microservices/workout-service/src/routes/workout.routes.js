const express = require('express');
const router = express.Router();
const { body, validationResult } = require('express-validator');
const Workout = require('../models/workout.model');
const { publishMessage } = require('../utils/rabbitmq');

// Middleware to validate workout data
const validateWorkout = [
  body('name').notEmpty().withMessage('Workout name is required'),
  body('type').isIn(['strength', 'cardio', 'flexibility', 'mixed']).withMessage('Invalid workout type'),
  body('duration').isNumeric().withMessage('Duration must be a number'),
  body('difficulty').isIn(['easy', 'medium', 'hard']).withMessage('Invalid difficulty level'),
  body('exercises').isArray().withMessage('Exercises must be an array')
];

// Create a new workout
router.post('/', validateWorkout, async (req, res) => {
  try {
    const errors = validationResult(req);
    if (!errors.isEmpty()) {
      return res.status(400).json({ errors: errors.array() });
    }

    const workout = new Workout(req.body);
    await workout.save();

    // Publish workout created event
    await publishMessage('workout.created', {
      userId: workout.userId,
      workoutId: workout._id,
      type: workout.type,
      date: workout.date
    });

    res.status(201).json(workout);
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
});

// Get all workouts for a user
router.get('/user/:userId', async (req, res) => {
  try {
    const workouts = await Workout.find({ userId: req.params.userId })
      .sort({ date: -1 });
    res.json(workouts);
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
});

// Get a specific workout
router.get('/:id', async (req, res) => {
  try {
    const workout = await Workout.findById(req.params.id);
    if (!workout) {
      return res.status(404).json({ message: 'Workout not found' });
    }
    res.json(workout);
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
});

// Update a workout
router.put('/:id', validateWorkout, async (req, res) => {
  try {
    const errors = validationResult(req);
    if (!errors.isEmpty()) {
      return res.status(400).json({ errors: errors.array() });
    }

    const workout = await Workout.findByIdAndUpdate(
      req.params.id,
      req.body,
      { new: true, runValidators: true }
    );

    if (!workout) {
      return res.status(404).json({ message: 'Workout not found' });
    }

    // Publish workout updated event
    await publishMessage('workout.updated', {
      userId: workout.userId,
      workoutId: workout._id,
      type: workout.type
    });

    res.json(workout);
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
});

// Delete a workout
router.delete('/:id', async (req, res) => {
  try {
    const workout = await Workout.findByIdAndDelete(req.params.id);
    if (!workout) {
      return res.status(404).json({ message: 'Workout not found' });
    }

    // Publish workout deleted event
    await publishMessage('workout.deleted', {
      userId: workout.userId,
      workoutId: workout._id
    });

    res.json({ message: 'Workout deleted successfully' });
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
});

module.exports = router; 