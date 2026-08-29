import 'dart:math';

class PresentValue {
  final num cashFlow;
  final num rate;
  final num period;
  const PresentValue(this.cashFlow, this.rate, this.period);
  @override
  bool operator ==(Object other) {
    identical(this, other) ||
        other is PresentValue &&
            runtimeType == other.runtimeType &&
            cashFlow == other.cashFlow &&
            rate == other.rate &&
            period == other.period;
    return super == other;
  }

  @override
  int get hashCode => Object.hash(cashFlow, rate, period);

  num calculate(num cashFlow, num rate, num period) {
    return cashFlow / (pow((1 + rate), period));
  }
}
