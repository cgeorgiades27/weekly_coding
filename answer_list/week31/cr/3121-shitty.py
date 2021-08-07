## First attempt. Only really works if it never encounters a dead end.

class ComboLock:

    def __init__(self, combo, xplist):
        self.combo = combo
        self.xplist = xplist
        self.scombo = str(self.combo)
        self.setxp = {str(x) for x in self.xplist}
        return

    def check_xp(self, com_check):
        if com_check not in self.xplist:
            return True
        else:
            return False

    def rotate(self):
        total = 0
        curr_combo = 0
        multi = 1
        for wheel in range(2, -1, -1):
            count = 0
            final = int(self.scombo[wheel])
            diff = 10 - final
            if final > 5:
                curr_combo += (multi * 10)
            while self.check_xp(curr_combo):
                if 5 >= final > count:
                    curr_combo += multi
                    count += 1
                elif final > 5 and count < diff:
                    curr_combo -= multi
                    count += 1
                else:
                    total += count
                    break
            multi *= 10
        return total, curr_combo



x = ComboLock(259, [112,999])
print(x.rotate())
