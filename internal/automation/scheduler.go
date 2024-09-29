package automation

import (
	"time"
)

// ScheduleTask представляет задачу, которую нужно выполнить по расписанию
type ScheduleTask struct {
	Interval time.Duration
	Action   func()
}

// Scheduler управляет задачами планировщика
type Scheduler struct {
	tasks []ScheduleTask
}

// NewScheduler создает новый планировщик
func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: make([]ScheduleTask, 0),
	}
}

// AddTask добавляет задачу в планировщик
func (s *Scheduler) AddTask(task ScheduleTask) {
	s.tasks = append(s.tasks, task)
}

// Start запускает выполнение задач планировщика
func (s *Scheduler) Start() {
	for _, task := range s.tasks {
		go func(t ScheduleTask) {
			ticker := time.NewTicker(t.Interval)
			defer ticker.Stop()
			for range ticker.C { // Используем for range для чтения из канала
				t.Action()
			}
		}(task)
	}
}
