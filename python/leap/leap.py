def is_leap_year(year):
    four, hundred, fourHundred = year % 4 == 0, year % 100 == 0, year % 400 == 0
    return True if fourHundred or (four and not hundred) else False
