package com.chainbase.udf.math;

import java.math.BigInteger;
import org.apache.flink.table.functions.ScalarFunction;

/**
 * Returns a BigInteger whose value is (number1 + number2).
 */
public class Add extends ScalarFunction {

  public String eval(String number1, String number2) {
    return new BigInteger(number1).add(new BigInteger(number2)).toString();
  }

}
