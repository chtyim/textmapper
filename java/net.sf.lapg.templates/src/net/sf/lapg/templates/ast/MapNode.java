package net.sf.lapg.templates.ast;

import java.util.Collection;
import java.util.HashMap;

import net.sf.lapg.templates.api.EvaluationContext;
import net.sf.lapg.templates.api.EvaluationException;
import net.sf.lapg.templates.api.IEvaluationEnvironment;

public class MapNode extends ExpressionNode {

	private static final String CONTEXTVAR = "context";
	private final ExpressionNode targetKey;
	private final ExpressionNode targetValue;
	private final ExpressionNode sourceList;

	public MapNode(ExpressionNode targetKey, ExpressionNode targetValue, ExpressionNode sourceList, String input, int line) {
		super(input, line);
		this.targetKey = targetKey;
		this.targetValue = targetValue;
		this.sourceList = sourceList;
	}

	@Override
	public Object evaluate(EvaluationContext context, IEvaluationEnvironment env) throws EvaluationException {
		Object select = env.evaluate(sourceList, context, false);
		Object prevVar = context.getVariable(CONTEXTVAR);
		context.setVariable(CONTEXTVAR, context);
		try {
			if (select instanceof Collection) {
				select = ((Collection<?>) select).toArray();
			}
			if (!(select instanceof Object[])) {
				throw new EvaluationException("`" + sourceList.toString() + "` should be array or collection");
			}
			Object[] source = (Object[]) select;

			if( targetKey == null ) {
				Object[] result = new Object[source.length];
				for( int i = 0; i < result.length; i++) {
					result[i] = env.evaluate(targetValue, new EvaluationContext(source[i],context), false);
				}
				return result;
			} else {
				HashMap<Object,Object> result = new HashMap<Object,Object>();
				for( int i = 0; i < source.length; i++) {
					result.put( env.evaluate(targetKey, new EvaluationContext(source[i],context), false), env.evaluate(targetValue, new EvaluationContext(source[i],context), false) );
				}
				return result;
			}
		} finally {
			context.setVariable(CONTEXTVAR, prevVar);
		}
	}

	@Override
	public void toString(StringBuffer sb) {
		sb.append("map(");
		if( targetKey != null ) {
			targetKey.toString(sb);
			sb.append("->");
		}
		targetValue.toString(sb);
		sb.append(") ");
		sourceList.toString(sb);
	}

}