export function dateToDateString(dateStr) {
  const date = new Date(dateStr)
  const day = date.getDate()
  let dayString
  if (~~(day / 10) === 0) {
      dayString = '0' + day
  } else {
      dayString = day.toString()
  }
  let monthString
  const month = date.getMonth()
  if (~~(month / 10) === 0) {
      monthString = '0' + month
  } else {
      monthString = month.toString()
  }
  const hours = date.getHours()
  let hoursString
  if (~~(hours / 10) === 0) {
      hoursString = `0${hours}`
  } else {
      hoursString = `${hours}`
  }
  const minutes = date.getMinutes()
  let minutesString
  if (~~(minutes / 10) === 0) {
      minutesString = `0${minutes}`
  } else {
      minutesString = `${minutes}`
  }
  const fullDate = `${dayString}.${monthString} в ${hoursString}:${minutesString}`
  return fullDate
}